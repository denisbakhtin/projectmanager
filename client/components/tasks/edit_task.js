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
        getProjects = () =>
        service.getProjects()
        .then((result) => projects = result.slice(0))
        .catch((error) => errors = responseErrors(error)),

        getProjectUsers = () =>
        service.getProjectUsers(task.project_id)
        .then((result) => pusers = result.slice(0))
        .catch((error) => errors = responseErrors(error)),

        get = () =>
        service.getTask(task.id)
        .then((result) => task = result)
        .catch((error) => errors = responseErrors(error)),

        getSteps = () =>
        service.getTaskSteps()
        .then((result) => steps = result.slice(0))
        .catch((error) => errors = responseErrors(error)),

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
            if (m.route.param('id')) {
                isNew = false
                task = {
                    id: m.route.param('id')
                }
                get()
            } else
                task = {}
            errors = []
            getSteps()
            getProjects().then(function(result) {
                if (!task.project_id && result.length) {
                    task.project_id = result[0].id
                    //is this correct?
                    getProjectUsers().then((res) => setProjectUserID(res.id))
                }
            })
        },
        view(vnode) {
            return m(".task", [
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
                            }, steps ?
                            steps.map((step) => {
                                return m('option', {
                                    value: step.id
                                }, step.name)
                            }) : null)
                    ),
                    m('.col-sm-5',
                        m('label', 'Project'),
                        m('select.form-control', {
                                onchange: function(e) {
                                    setProjectID(e.target.value)
                                    getProjectUsers()
                                },
                                value: task.project_id
                            }, projects ?
                            projects.map((proj) => {
                                return m('option', {
                                    value: proj.id
                                }, proj.name)
                            }) : null)
                    ),
                    m('.col-sm-5',
                        m('label', 'Assigned user'),
                        m('select.form-control', {
                                onchange: (e) => setProjectUserID(e.target.value),
                                value: task.project_user_id
                            }, pusers ?
                            pusers.map((puser) => {
                                return m('option', {
                                    value: puser.id
                                }, puser.user.name + "(" + puser.role.name + ")")
                            }) : null)
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
                    (isNew || task.files != undefined) ? m(files, {
                        files: task.files,
                        onchange: setFiles
                    }) : null,
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
            ])
        }
    }
}
