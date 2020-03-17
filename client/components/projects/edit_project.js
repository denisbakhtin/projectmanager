import m from 'mithril'
import error from '../shared/error'
import files from '../attached_files/files'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors,
} from '../../utils/helpers'
import Auth from '../../utils/auth'
import service from '../../utils/service.js'
//import pusers from '../project_users/pusers.js'

export default function Project() {
    let errors = [],
        project = {},
        categories = [],
        isNew = true,
        loaded = false,

        setName = (name) => project.name = name,
        setCategoryID = (id) => project.category_id = id,
        setDescription = (description) => project.description = description,
        //setProjectUsers = (pu) => project.project_users = pu,
        setFiles = (files) => project.files = files,

        getOwnerName = () => (project.owner || Auth.getAuthenticatedUser()).name,

        //requests
        newProject = () =>
            service.newProject()
                .then((result) => {
                    project = result.project
                    project.files = []
                    categories = result.categories
                    loaded = true
                }).catch((error) => errors = responseErrors(error)),

        editProject = (id) =>
            service.editProject(id)
                .then((result) => {
                    project = result.project
                    categories = result.categories
                    loaded = true
                }).catch((error) => errors = responseErrors(error)),

        create = () =>
            service.createProject(project)
                .then((result) => {
                    addSuccess("Project created.")
                    m.route.set('/projects')
                })
                .catch((error) => errors = responseErrors(error)),

        update = () =>
            service.updateProject(project.id, project)
                .then((result) => {
                    addSuccess("Project updated.")
                    window.history.back()
                })
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            isNew = (m.route.param('id') == undefined)
            if (isNew)
                newProject()
            else
                editProject(m.route.param('id'))
        },

        view(vnode) {
            return m(".projects", (loaded) ? [
                m('h1.title', (isNew) ? 'New project' : 'Edit project'),
                m('.form-group', [
                    m('label', 'Project name'),
                    m('input.form-control[type=text]', {
                        oncreate: (el) => el.dom.focus(),
                        oninput: (e) => setName(e.target.value),
                        placeholder: 'e.g. Web-site development',
                        value: project.name
                    })
                ]),
                (categories.length > 0) ?
                    m('.form-group.form-row', [
                        m('.col-auto', [
                            m('label', 'Category'),
                            m('select.form-control', {
                                onchange: (e) => setCategoryID(e.target.value),
                                value: project.category_id ?? ""
                            }, [
                                m('option[value=""][disabled=disabled][hidden=hidden]', "Select category..."),
                                categories.map((cat) => {
                                    return m('option', {
                                        value: cat.id,
                                        selected: (cat.id == project.category_id)
                                    }, cat.name)
                                }),
                            ])
                        ]),
                    ]) : null,
                /*
                    m('.form-group', [
                    m('label', "Assigned users"),
                    m(pusers, {
                        project_id: project.id,
                        project_users: project.project_users,
                        onchange: setProjectUsers
                    }),
                ]),
                */
                m('.form-group', [
                    m('label', "Description (supports Markdown)"),
                    m('textarea.form-control', {
                        oninput: (e) => setDescription(e.target.value),
                        value: project.description
                    })
                ]),
                m('.form-group', [
                    m('label', "Attached files"),
                    m(files, {
                        files: project.files,
                        onchange: setFiles
                    }),
                ]),
                m('.mb-2', m(error, { errors: errors })),
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
