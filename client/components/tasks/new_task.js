import m from 'mithril'
import state from './state'
import form from './form'

const Task = {
    oninit(vnode) {
        state.task = {}
        state.errors = []
    },
    
    view(vnode) {
        let ui = vnode.state
        return m(".task", [
            m('h1.mb-4', 'New task'),
            form(),
            m('.actions', [
                m('button.btn.btn-primary.mr-2[type=button]', { onclick: state.create }, "Create"),
                m('button.btn.btn-secondary[type=button]', { onclick: () => { m.route.set('/tasks') } }, "Cancel")
            ]),
        ])
    }
}

export default Task;