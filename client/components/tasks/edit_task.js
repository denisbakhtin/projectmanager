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
        errors = [],
        steps = [],
        projects = [],
        isNew = true,
        loaded = false,

        setName = (name) => task.name = name,
        setDescription = (description) => task.description = description,
        setStepID = (step_id) => task.task_step_id = step_id,
        setProjectUserID = (project_user_id) => task.project_user_id = project_user_id,
        setProjectID = (project_id) => task.project_id = project_id,
        setFiles = (files) => task.files = files,
        validate = () => {
            errors = []
            if (!task.name)
                errors.push("Task name is required.")
            return errors.length == 0
        },

        //requests
        newTask = () =>
        service.newTask()
        .then((result) => {
            projects = result.projects
            task = result.task
            steps = result.task_steps
            loaded = true;
        }).catch((error) => errors = responseErrors(error)),

        editTask = (id) =>
        service.editTask(id)
        .then((result) => {
            projects = result.projects
            task = result.task
            steps = result.task_steps
            loaded = true;
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
            m.route.set('/tasks')
        })
        .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            isNew = (m.route.param('id') == undefined)
            if (isNew)
                newTask()
            else
                editTask(m.route.param('id'))
        },
        view(vnode) {
            return m(".task", (loaded) ? [
                m('h1.mb-4', (isNew) ? 'New task' : 'Edit task'),
                m('.form-group', [
                    m('label', 'Task name'),
                    m('input.form-control[type=text]', {
                        oncreate: (el) => {
                            el.dom.focus()
                        },
                        oninput: (e) => setName(e.target.value),
                        value: task.name
                    })
                ]),
                m('.form-group.form-row', [
                    m('.col-sm-2',
                        m('label', 'State'),
                        m('select.form-control', {
                            onchange: (e) => setStepID(e.target.value),
                            value: task.task_step_id
                        },
                            steps.map((step) => {
                                return m('option', {
                                    value: step.id,
                                    selected: (step.id == task.task_step_id)
                                }, step.name)
                            })
                        )
                    ),
                    m('.col-sm-5',
                        m('label', 'Project'),
                        m('select.form-control', {
                            onchange: function(e) {
                                setProjectID(e.target.value)
                                getProjectUsers()
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
                ]),
                m('.form-group', [
                    m('label', "Description"),
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
                    }, "Save"),
                    m('button.btn.btn-secondary[type=button]', {
                        onclick: () => {
                            m.route.set('/tasks')
                        }
                    }, "Cancel")
                ]),
            ] : m('Loading...'))
        }
    }
}
