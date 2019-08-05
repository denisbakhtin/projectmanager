import m from 'mithril'
import state from './state'
import form from './form'

const TaskStep = {
    oninit(vnode) {
        state.step = {}
        state.errors = []
    },
    
    view(vnode) {
        let ui = vnode.state
        return m(".task_step", [
            m('h1.mb-4', 'New task step'),
            form(),
            m('.actions', [
                m('button.btn.btn-primary.mr-2[type=button]', { onclick: state.create }, "Create"),
                m('button.btn.btn-secondary[type=button]', { onclick: () => { m.route.set('/task_steps') } }, "Cancel")
            ]),
        ])
    }
}

export default TaskStep;