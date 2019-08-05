import m from 'mithril'
import state from './state'

const Status = {
    oninit(vnode) {
        state.errors = []
        state.status = { id: m.route.param('id') }
        state.get()
    },
    
    view(vnode) {
        let ui = vnode.state
        return m(".status", [
            state.status.name ? [
                m('h1.mb-2', state.status.name),
                state.status.description ? [
                    m('h3', "Description"),
                    m('p', state.status.description)
                ] : null
            ] : null, 
            m('.actions', [
                m('button.btn.btn-primary.mr-2[type=button]', { onclick: () => { m.route.set('/statuses/edit/' + state.status.id) } }, "Edit"),
                m('button.btn.btn-secondary.mr-2[type=button]', { onclick: () => { m.route.set('/statuses') } }, "Back to list"),
                m('button.btn.btn-outline-danger[type=button]', { onclick: state.destroy }, "Remove status")
            ])
        ])
    }
}

export default Status;