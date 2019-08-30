import m from 'mithril'
import form from './form'
import state from './state'

const Status = {
    oninit(vnode) {
        state.errors = []
        state.status = { id: m.route.param('id')}
        state.get()
    },
    
    view(vnode) {
        let ui = vnode.state
        return m(".projects", [
            m('h1.mb-4', 'Edit status'),
            form(),
            m('.actions', [
                m('button.btn.btn-primary.mr-2[type=button]', { onclick: state.update }, "Update"),
                m('button.btn.btn-secondary[type=button]', { onclick: () => { window.history.back() } }, "Cancel")
            ]),
        ])
    }
}

export default Status;