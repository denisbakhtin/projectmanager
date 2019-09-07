import m from 'mithril'
import { ISODateToHtml5 } from '../../utils/helpers'
import error from '../shared/error'
import service from '../../utils/service.js'

export default function Projects() {
    var projects = [],
        errors = []

    function getAll() {
        return service.getProjects()
            .then((result) => projects = result.map((r) => fromGo(r)))
            .catch((error) => errors = responseErrors(error))
    }
    function fromGo(proj) {
        return Object.assign(proj, {
            start_date: ISODateToHtml5(proj.start_date || null),
            end_date: ISODateToHtml5(proj.end_date || null)
        })
    }

    return {
        oninit(vnode) {
            errors = []
            getAll()
        },

        view(vnode) {
            return m(".projects", [
                m('h1.mb-4', 'Projects'),
                projects ? projects.map((proj) => {
                    return m('.card.mb-2', { onclick: () => { m.route.set('/projects/' + proj.id) } }, [
                        m('.card-body', [
                            m('h5.card-title', proj.name),
                            m('h6.card-subtitle.mb-2.text-muted', [
                                proj.start_date ? [
                                    m('span', 'Starts: '),
                                    m('span', ISODateToHtml5(proj.start_date, '-')),
                                ] : null,
                                proj.end_date ? [
                                    m('span', ' Ends: '),
                                    m('span', ISODateToHtml5(proj.end_date, '-')),
                                ] : null,
                                m('span', ' Owner: '),
                                m('span', (proj.owner && proj.owner.name) ? proj.owner.name : "-"),
                                m('span', ' Status: '),
                                m('span', proj.status.name)
                            ]),
                            m('p.card-text', proj.description)
                        ])
                    ])
                }) : null,
                m('.actions.mt-4', [
                    m('button.btn.btn-primary[type=button]', { onclick: () => { m.route.set('/projects/new') } }, "New project")
                ]),
            ])
        }
    }
}
