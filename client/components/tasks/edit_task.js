import m from 'mithril'
import service from '../../utils/service.js'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors,
    isZeroDate
} from '../../utils/helpers'
import files from '../attached_files/files'
import error from '../shared/error'
import moment from 'moment'
const format = 'YYYY-MM-DD'

//binary week days constants
const WeekdayMasks = Object.freeze({
    1: 1,
    2: 2,
    3: 4,
    4: 8,
    5: 16,
    6: 32,
    7: 64
})

export default function Task() {
    let task = {},
        pusers = [],
        project_id,
        errors = [],
        categories = [],
        projects = [],
        isNew = true,
        loaded = false,

        setName = (name) => task.name = name,
        setDescription = (description) => task.description = description,
        setCategoryID = (id) => task.category_id = id,
        //setProjectUserID = (project_user_id) => task.project_user_id = project_user_id,
        setProjectID = (project_id) => task.project_id = project_id,
        setFiles = (files) => task.files = files,
        setPriority = (prio) => task.priority = prio,
        getStartDate = () => (!isZeroDate(task.start_date)) ? moment(task.start_date).format(format) : undefined,
        setStartDate = (date) => task.start_date = (date) ? (new Date(date)).toISOString() : null,
        getEndDate = () => (!isZeroDate(task.end_date)) ? moment(task.end_date).format(format) : undefined,
        setEndDate = (date) => task.end_date = (date) ? (new Date(date)).toISOString() : null,
        setPeriodicity = (val) => {
            task.periodicity.periodicity_type = val
            if (val == 3 || val == 4) {
                task.periodicity.repeating_from = task.start_date
                task.periodicity.repeating_to = task.end_date
            }
        },
        toggleWeekday = (day) => task.periodicity.weekdays = task.periodicity.weekdays ^ WeekdayMasks[day],
        getWeekday = (day) => task.periodicity.weekdays & WeekdayMasks[day],
        getRepeatingFrom = () => (!isZeroDate(task.periodicity.repeating_from)) ? moment(task.periodicity.repeating_from).format(format) : undefined,
        setRepeatingFrom = (date) => task.periodicity.repeating_from = (date) ? (new Date(date)).toISOString() : null,
        getRepeatingTo = () => (!isZeroDate(task.periodicity.repeating_to)) ? moment(task.periodicity.repeating_to).format(format) : undefined,
        setRepeatingTo = (date) => task.periodicity.repeating_to = (date) ? (new Date(date)).toISOString() : null,
        //need to do something with timezone later... browser uses localtime always

        //requests
        newTask = () =>
            service.newTask(project_id)
                .then((result) => {
                    projects = result.projects
                    categories = result.categories
                    task = result.task
                    loaded = true
                }).catch((error) => errors = responseErrors(error)),

        editTask = (id) =>
            service.editTask(id)
                .then((result) => {
                    projects = result.projects
                    categories = result.categories
                    task = result.task
                    loaded = true
                }).catch((error) => errors = responseErrors(error)),

        create = () =>
            service.createTask(task)
                .then((result) => {
                    addSuccess("Task created.")
                    window.history.back()
                })
                .catch((error) => errors = responseErrors(error)),

        update = () => {
            service.updateTask(task.id, task)
                .then((result) => {
                    addSuccess("Task updated.")
                    window.history.back();
                })
                .catch((error) => errors = responseErrors(error))
        }

    /* getProjectUsers = () =>
        service.getProjectUsers(task.project_id)
            .then((result) => pusers = result)
            .catch((error) => errors = responseErrors(error))
    */

    return {
        oninit(vnode) {
            project_id = m.route.param('project_id')
            isNew = (m.route.param('id') == undefined)
            if (isNew)
                newTask()
            else
                editTask(m.route.param('id'))
        },
        view(vnode) {
            return m(".task", (loaded) ? [
                m('h1.title', (isNew) ? 'New task' : 'Edit task'),
                m('.form-group', [
                    m('label', 'Task name'),
                    m('input.form-control[type=text]', {
                        oncreate: (el) => {
                            el.dom.focus()
                        },
                        oninput: (e) => setName(e.target.value),
                        placeholder: "e.g. Post a new article",
                        value: task.name
                    })
                ]),
                m('.form-group.form-row', [
                    m('.col-sm-5',
                        m('label', 'Project'),
                        m('select.form-control', {
                            onchange: function (e) {
                                setProjectID(e.target.value)
                                //getProjectUsers()
                            },
                            value: task.project_id
                        },
                            projects.map((proj) => {
                                return m('option', {
                                    value: proj.id,
                                    selected: (proj.id == task.project_id)
                                }, proj.name)
                            })
                        )
                    ),
                    (categories.length > 0) ?
                        m('.col-sm-5', [
                            m('label', 'Category'),
                            m('select.form-control', {
                                onchange: (e) => setCategoryID(e.target.value),
                                value: task.category_id ?? ""
                            }, [
                                m('option[value=""][disabled=disabled][hidden=hidden]', "Select category..."),
                                categories.map((cat) => {
                                    return m('option', {
                                        value: cat.id,
                                        selected: (cat.id == task.category_id)
                                    }, cat.name)
                                }),
                            ])
                        ]) : null,
                    /*
                    m('.col-sm-5',
                        m('label', 'Assigned user'),
                        m('select.form-control', {
                            onchange: (e) => setProjectUserID(e.target.value),
                            value: task.project_user_id
                        },
                            pusers.map((puser) => {
                                return m('option', {
                                    value: puser.id,
                                    selected: (puser.id == task.project_user_id)
                                }, puser.user.name + "(" + puser.role.name + ")")
                            })
                        )
                    ),
                    */
                ]),

                //priority
                m('.form-group', [
                    m('.custom-control.custom-radio.custom-control-inline', [
                        m('input#priority1.custom-control-input[type=radio][name=priority][value=1]', {
                            checked: task.priority == 1,
                            onchange: (e) => setPriority(e.target.value)
                        }),
                        m('label.custom-control-label.mr-1[for=priority1]', 'Do Now'),
                    ]),
                    m('.custom-control.custom-radio.custom-control-inline', [
                        m('input#priority2.custom-control-input[type=radio][name=priority][value=2]', {
                            checked: task.priority == 2,
                            onchange: (e) => setPriority(e.target.value)
                        }),
                        m('label.custom-control-label.mr-1[for=priority2]', 'Do Next'),
                    ]),
                    m('.custom-control.custom-radio.custom-control-inline', [
                        m('input#priority3.custom-control-input[type=radio][name=priority][value=3]', {
                            checked: task.priority == 3,
                            onchange: (e) => setPriority(e.target.value)
                        }),
                        m('label.custom-control-label.mr-1[for=priority3]', 'Do Later'),
                    ]),
                    m('.custom-control.custom-radio.custom-control-inline', [
                        m('input#priority4.custom-control-input[type=radio][name=priority][value=4]', {
                            checked: task.priority == 4,
                            onchange: (e) => setPriority(e.target.value)
                        }),
                        m('label.custom-control-label.mr-1[for=priority4]', 'Do Last'),
                    ]),
                ]), //end priority

                m('.form-group.form-row', [
                    m('.col-auto', [
                        m('label', "Start date"),
                        m('input[type=date].form-control', {
                            oninput: (e) => setStartDate(e.target.value),
                            value: getStartDate()
                        })
                    ]),
                    m('.col-auto', [
                        m('label', "End date"),
                        m('input[type=date].form-control', {
                            oninput: (e) => setEndDate(e.target.value),
                            value: getEndDate()
                        })
                    ])
                ]),

                //periodicity
                m('.form-group', m('div', [
                    m('.custom-control.custom-radio.custom-control-inline', [
                        m('input#no-repeat.custom-control-input[type=radio][name=periodicity][value=0]', {
                            checked: task.periodicity.periodicity_type == 0,
                            onchange: (e) => setPeriodicity(e.target.value)
                        }),
                        m('label.custom-control-label.mr-1[for=no-repeat]', "Do not repeat"),
                    ]),
                    m('.custom-control.custom-radio.custom-control-inline', [
                        m('input#daily.custom-control-input[type=radio][name=periodicity][value=1]', {
                            checked: task.periodicity.periodicity_type == 1,
                            onchange: (e) => setPeriodicity(e.target.value)
                        }),
                        m('label.custom-control-label.mr-1[for=daily]', 'Repeat Daily'),
                    ]),
                    m('.custom-control.custom-radio.custom-control-inline', [
                        m('input#weekly.custom-control-input[type=radio][name=periodicity][value=2]', {
                            checked: task.periodicity.periodicity_type == 2,
                            onchange: (e) => setPeriodicity(e.target.value)
                        }),
                        m('label.custom-control-label.mr-1[for=weekly]', 'Weekly'),
                    ]),
                    m('.custom-control.custom-radio.custom-control-inline', [
                        m('input#monthly.custom-control-input[type=radio][name=periodicity][value=3]', {
                            checked: task.periodicity.periodicity_type == 3,
                            onchange: (e) => setPeriodicity(e.target.value)
                        }),
                        m('label.custom-control-label.mr-1[for=monthly]', 'Monthly'),
                    ]),
                    m('.custom-control.custom-radio.custom-control-inline', [
                        m('input#yearly.custom-control-input[type=radio][name=periodicity][value=4]', {
                            checked: task.periodicity.periodicity_type == 4,
                            onchange: (e) => setPeriodicity(e.target.value)
                        }),
                        m('label.custom-control-label.mr-1[for=yearly]', 'Yearly'),
                    ]),

                    (task.periodicity.periodicity_type == 2) ? m('.weekdays.mt-2', [
                        m('.custom-control.custom-checkbox.custom-control-inline', [
                            m('input.custom-control-input#monday[type=checkbox]', {
                                checked: getWeekday(1),
                                onchange: (e) => toggleWeekday(1)
                            }),
                            m('label.custom-control-label[for=monday]', 'Mon')
                        ]),
                        m('.custom-control.custom-checkbox.custom-control-inline', [
                            m('input.custom-control-input#tuesday[type=checkbox]', {
                                checked: getWeekday(2),
                                onchange: (e) => toggleWeekday(2)
                            }),
                            m('label.custom-control-label[for=tuesday]', 'Tue')
                        ]),
                        m('.custom-control.custom-checkbox.custom-control-inline', [
                            m('input.custom-control-input#wednesday[type=checkbox]', {
                                checked: getWeekday(3),
                                onchange: (e) => toggleWeekday(3)
                            }),
                            m('label.custom-control-label[for=wednesday]', 'Wed')
                        ]),
                        m('.custom-control.custom-checkbox.custom-control-inline', [
                            m('input.custom-control-input#thursday[type=checkbox]', {
                                checked: getWeekday(4),
                                onchange: (e) => toggleWeekday(4)
                            }),
                            m('label.custom-control-label[for=thursday]', 'Thu')
                        ]),
                        m('.custom-control.custom-checkbox.custom-control-inline', [
                            m('input.custom-control-input#friday[type=checkbox]', {
                                checked: getWeekday(5),
                                onchange: (e) => toggleWeekday(5)
                            }),
                            m('label.custom-control-label[for=friday]', 'Fri')
                        ]),
                        m('.custom-control.custom-checkbox.custom-control-inline', [
                            m('input.custom-control-input#satursday[type=checkbox]', {
                                checked: getWeekday(6),
                                onchange: (e) => toggleWeekday(6)
                            }),
                            m('label.custom-control-label[for=satursday]', 'Sat')
                        ]),
                        m('.custom-control.custom-checkbox.custom-control-inline', [
                            m('input.custom-control-input#sunday[type=checkbox]', {
                                checked: getWeekday(7),
                                onchange: (e) => toggleWeekday(7)
                            }),
                            m('label.custom-control-label[for=sunday]', 'Sun')
                        ]),
                    ]) : null,

                    (task.periodicity.periodicity_type == 3 || task.periodicity.periodicity_type == 4) ? m('.form-group.form-row.mt-2', [
                        m('.col-auto', [
                            m('label', "Starting from"),
                            m('input[type=date].form-control', {
                                oninput: (e) => setRepeatingFrom(e.target.value),
                                value: getRepeatingFrom()
                            })
                        ]),
                        m('.col-auto', [
                            m('label', "Ending at"),
                            m('input[type=date].form-control', {
                                oninput: (e) => setRepeatingTo(e.target.value),
                                value: getRepeatingTo()
                            })
                        ])
                    ]) : null,
                ])), //end periodicity

                m('.form-group', [
                    m('label', "Description (supports Markdown)"),
                    m('textarea.form-control', {
                        oninput: (e) => setDescription(e.target.value),
                        value: task.description
                    })
                ]),
                m('.form-group', [
                    m('label', "Attached files"),
                    m(files, {
                        files: task.files,
                        onchange: setFiles
                    }),
                ]),
                m('.mb-2', m(error, {
                    errors: errors
                })),

                m('.actions', [
                    m('button.btn.btn-primary.mr-2[type=button]', {
                        onclick: (isNew) ? create : update
                    }, [
                        m("i.fa.fa-check.mr-1"),
                        "Submit"
                    ]),
                    m('button.btn.btn-outline-secondary[type=button]', {
                        onclick: () => window.history.back()
                    }, "Cancel")
                ]),
            ] : m('Loading...'))
        }
    }
}
