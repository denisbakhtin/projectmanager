import m from 'mithril'
import error from '../shared/error'
import state from './state'

const Project = {
    oninit(vnode) {
        state.errors = []
        state.project = { id: m.route.param('id') }
        state.get()
    },
    
    view(vnode) {
        let ui = vnode.state
        return m(".project", [
            state.project.name ? [
                m('h1.mb-2', state.project.name),
                m('h6.mb-4.text-muted', [
                    m('span', 'Starts: '),
                    m('span', (state.project.start_date) ? new Date(state.project.start_date).toISOString().slice(0, 10) : "-"),
                    m('span', ' Ends: '),
                    m('span', (state.project.end_date) ? new Date(state.project.end_date).toISOString().slice(0, 10) : "-"),
                    m('span', ' Owner: '),
                    m('span', (state.project.owner && state.project.owner.name) ? state.project.owner.name : "-"),
                    m('span', ' Status: '),
                    m('span', state.project.status)
                ]),
                state.project.description ? [
                    m('h3', "Description"),
                    m('p', state.project.description)
                ] : null
            ] : null, 
            m('.actions', [
                m('button.btn.btn-primary.mr-2[type=button]', { onclick: () => { m.route.set('/projects/edit/' + state.project.id) } }, "Edit"),
                m('button.btn.btn-secondary.mr-2[type=button]', { onclick: () => { m.route.set('/projects') } }, "Back to list"),
                m('button.btn.btn-outline-danger[type=button]', { onclick: state.destroy }, "Remove project")
            ])
        ])
    }
}

export default Project;