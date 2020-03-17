import m from 'mithril'
import service from '../../utils/service.js'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors
} from '../../utils/helpers'
import files from '../attached_files/files'
import error from '../shared/error'

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

        //requests
        newTask = () =>
            service.newTask()
                .then((result) => {
                    projects = result.projects
                    categories = result.categories
                    task = result.task
                    task.project_id = project_id ?? ((projects && projects.length > 0) ? projects[0].id : null)
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
                    m.route.set('/tasks')
                })
                .catch((error) => errors = responseErrors(error)),

        update = () =>
            service.updateTask(task.id, task)
                .then((result) => {
                    addSuccess("Task updated.")
                    window.history.back();
                })
                .catch((error) => errors = responseErrors(error))
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
                ]),
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
