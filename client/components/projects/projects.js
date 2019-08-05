import m from 'mithril'
import { ISODateToHtml5 } from '../../utils/helpers'
import error from '../shared/error'
import state from './state'

const Projects = {
    oninit(vnode) {
        state.errors = []
        state.getAll()
    },
    
    view(vnode) {
        let ui = vnode.state
        return m(".projects", [
            m('h1.mb-4', 'Projects'),
            state.projects ? state.projects.map((proj) => {
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

export default Projects;