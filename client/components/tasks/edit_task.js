import m from 'mithril'
import form from './form'
import state from './state'

const Task = {
    oninit(vnode) {
        state.errors = []
        state.task = { id: m.route.param('id')}
        state.get()
    },
    
    view(vnode) {
        let ui = vnode.state
        return m(".task", [
            m('h1.mb-4', 'Edit task'),
            form(),
            m('.actions', [
                m('button.btn.btn-primary.mr-2[type=button]', { onclick: state.update }, "Update"),
                m('button.btn.btn-secondary[type=button]', { onclick: () => { window.history.back() } }, "Cancel")
            ]),
        ])
    }
}

export default Task;