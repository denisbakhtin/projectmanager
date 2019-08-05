import m from 'mithril'
import state from './state'
import form from './form'

const Status = {
    oninit(vnode) {
        state.status = {}
        state.errors = []
    },
    
    view(vnode) {
        let ui = vnode.state
        return m(".status", [
            m('h1.mb-4', 'New status'),
            form(),
            m('.actions', [
                m('button.btn.btn-primary.mr-2[type=button]', { onclick: state.create }, "Create"),
                m('button.btn.btn-secondary[type=button]', { onclick: () => { m.route.set('/statuses') } }, "Cancel")
            ]),
        ])
    }
}

export default Status;