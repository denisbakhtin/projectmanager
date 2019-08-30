import m from 'mithril'
import error from '../shared/error'
import state from './state'
import form from './form'
import statuses_state from '../statuses/state'

const Project = {
    oninit(vnode) {
        state.errors = []
        state.project = { id: m.route.param('id')}
        state.get()
        statuses_state.getAll()
    },
    
    view(vnode) {
        let ui = vnode.state
        return m(".projects", [
            m('h1.mb-4', 'Edit project'),
            state.project.name ? form() : null,
            m('.actions', [
                m('button.btn.btn-primary.mr-2[type=button]', { onclick: state.update }, "Update"),
                m('button.btn.btn-secondary[type=button]', { onclick: () => { window.history.back() } }, "Cancel")
            ]),
        ])
    }
}

export default Project;