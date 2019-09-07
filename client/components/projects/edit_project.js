import m from 'mithril'
import error from '../shared/error'
import files from '../attached_files/files'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors,
    ISODateToHtml5
} from '../../utils/helpers'
import Auth from '../../utils/auth'
import service from '../../utils/service.js'
import pusers from '../project_users/pusers.js'

export default function Project() {
    let errors = [],
        project = {},
        isNew = true,
        statuses = [],

        setName = (name) => project.name = name,
        setDescription = (description) => project.description = description,
        setStatusId = (status_id) => project.status_id = status_id,
        setStartDate = (date) => {
            project.start_date = date
            if (date && project.end_date && project.end_date < date)
                project.end_date = null
        },
        setEndDate = (date) => {
            project.end_date = date
            if (date && project.start_date && project.start_date > date)
                project.start_date = null
        },
        setProjectUsers = (pusers) => project.project_users = pusers,
        setFiles = (files) => project.files = files,

        validate = () => {
            errors = []
            if (!project.name)
                errors.push("Project name is required.")
            if (project.start_date && project.end_date && project.start_date > project.end_date)
                errors.push("End date cannot be earlier than start date.")
            return errors.length == 0
        },

        toGo = (proj) => {
            let obj = Object.assign(proj, {
                status_id: proj.status_id ? "" + proj.status_id : undefined,
                start_date: proj.start_date ? new Date(proj.start_date).toISOString() : undefined,
                end_date: proj.end_date ? new Date(proj.end_date).toISOString() : undefined,
                project_users: (proj.project_users) ? proj.project_users.map((pu) => {
                    pu.role_id = "" + pu.role_id;
                    return pu
                }) : undefined,
                files: proj.files
            })
            return obj
        },

        fromGo = (proj) =>
        Object.assign(proj, {
            start_date: ISODateToHtml5(proj.start_date || null),
            end_date: ISODateToHtml5(proj.end_date || null)
        }),

        getOwnerName = () => (project.owner || Auth.getAuthenticatedUser()).name,

        //requests
        get = () =>
        service.getProject(project.id)
        .then((result) => project = fromGo(result))
        .catch((error) => errors = responseErrors(error)),

        getStatuses = () =>
        service.getStatuses()
        .then((result) => statuses = result.slice(0))
        .catch((error) => errors = responseErrors(error)),

        create = () =>
        service.createProject(toGo(project))
        .then((result) => {
            addSuccess("Project created.")
            m.route.set('/projects')
        })
        .catch((error) => errors = responseErrors(error)),

        update = () =>
        service.updateProject(project.id, toGo(project))
        .then((result) => {
            addSuccess("Project updated.")
            m.route.set('/projects')
        })
        .catch((error) => errors = responseErrors(error)),

        destroy = () =>
        service.deleteProject(project.id)
        .then((result) => {
            addSuccess("Project removed.")
            m.route.set('/projects', {}, {
                replace: true
            })
        })
        .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            if (m.route.param('id')) {
                isNew = false
                project = {
                    id: m.route.param('id')
                }
                get()
            } else {
                project = {}
            }
            errors = []
            getStatuses().then((result) => {
                if (!project.status_id)
                    project.status_id = result[0].id
            })
        },

        view(vnode) {
            return m(".projects", [
                m('h1.mb-4', (isNew) ? 'New project' : 'Edit project'),
                m('.form-group', [
                    m('label', 'Project name'),
                    m('input.form-control[type=text]', {
                        oncreate: (el) => {
                            el.dom.focus()
                        },
                        oninput: (e) => setName(e.target.value),
                        value: project.name
                    })
                ]),
                m('.form-inline', [
                    m('.form-group.mb-3.mr-4', [
                        m('label.mr-2', 'Start date'),
                        m('input.form-control[type=date]', {
                            oninput: (e) => setStartDate(e.target.value),
                            value: project.start_date
                        })
                    ]),
                    m('.form-group.mb-3.mr-4', [
                        m('label.mr-2', 'End date'),
                        m('input.form-control[type=date]', {
                            oninput: (e) => setEndDate(e.target.value),
                            min: project.start_date,
                            value: project.end_date
                        })
                    ]),
                    m('.form-group.mb-3.mr-4', [
                        m('label.mr-2', 'Owner'),
                        m('input.form-control[type=text][disabled]', {
                            value: getOwnerName()
                        })
                    ]),
                    m('.form-group.mb-3', [
                        m('label.mr-2', 'Status'),
                        m('select.form-control', {
                                onchange: (e) => setStatusId(e.target.value),
                                value: project.status_id
                            }, statuses ?
                            statuses.map((status) => {
                                return m('option', {
                                    value: status.id
                                }, status.name)
                            }) : null)
                    ]),
                ]),
                m('.form-group', [
                    m('label', "Assigned users"),
                    m(pusers, {
                        project_id: project.id,
                        project_users: project.project_users,
                        onchange: setProjectUsers
                    }),
                ]),
                m('.form-group', [
                    m('label', "Description"),
                    m('textarea.form-control', {
                        oninput: (e) => setDescription(e.target.value),
                        value: project.description
                    })
                ]),
                m('.form-group', [
                    m('label', "Attached files"),
                    (isNew || project.files != undefined) ? m(files, {
                        files: project.files,
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
                            window.history.back()
                        }
                    }, "Cancel")
                ]),
            ])
        }
    }
}
