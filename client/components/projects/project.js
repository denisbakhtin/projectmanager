import m from 'mithril'
import error from '../shared/error'
import service from '../../utils/service.js'
import {
    responseErrors,
    ISODateToHtml5
} from '../../utils/helpers'

export default function Project() {
    let project = {},
        errors = [],

        get = () =>
        service.getProject(project.id)
        .then((result) => project = fromGo(result))
        .catch((error) => errors = responseErrors(error)),

        fromGo = (proj) =>
        Object.assign(proj, {
            start_date: ISODateToHtml5(proj.start_date || null),
            end_date: ISODateToHtml5(proj.end_date || null)
        }),

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
            project = {
                id: m.route.param('id')
            }
            get()
        },

        view(vnode) {
            return m(".project", [
                project.name ? [
                    m('h1.mb-2', project.name),
                    m('h6.mb-4.text-muted', [
                        m('span', 'Starts: '),
                        m('span', (project.start_date) ? new Date(project.start_date).toISOString().slice(0, 10) : "-"),
                        m('span', ' Ends: '),
                        m('span', (project.end_date) ? new Date(project.end_date).toISOString().slice(0, 10) : "-"),
                        m('span', ' Owner: '),
                        m('span', (project.owner && project.owner.name) ? project.owner.name : "-"),
                        m('span', ' Status: '),
                        m('span', project.status)
                    ]),
                    project.description ? [
                        m('h3', "Description"),
                        m('p', project.description)
                    ] : null
                ] : null,
                m('.actions', [
                    m('button.btn.btn-primary.mr-2[type=button]', {
                        onclick: () => {
                            m.route.set('/projects/edit/' + project.id)
                        }
                    }, "Edit"),
                    m('button.btn.btn-secondary.mr-2[type=button]', {
                        onclick: () => {
                            m.route.set('/projects')
                        }
                    }, "Back to list"),
                    m('button.btn.btn-outline-danger[type=button]', {
                        onclick: destroy
                    }, "Remove project")
                ])
            ])
        }
    }
}
