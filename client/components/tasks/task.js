import m from 'mithril'
import state from './state'

const TaskStep = {
    oninit(vnode) {
        state.errors = []
        state.step = { id: m.route.param('id') }
        state.get()
    },
    
    view(vnode) {
        let ui = vnode.state
        return m(".task_step", [
            state.step.name ? [
                m('h1.mb-2', state.step.name),
                m('p', 'Is final: ' + state.step.is_final)
            ] : null, 
            m('.actions', [
                m('button.btn.btn-primary.mr-2[type=button]', { onclick: () => { m.route.set('/task_steps/edit/' + state.step.id) } }, "Edit"),
                m('button.btn.btn-secondary.mr-2[type=button]', { onclick: () => { m.route.set('/task_steps') } }, "Back to list"),
                m('button.btn.btn-outline-danger[type=button]', { onclick: state.destroy }, "Remove step")
            ])
        ])
    }
}

export default TaskStep;