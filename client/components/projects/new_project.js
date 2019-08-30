import m from 'mithril'
import error from '../shared/error'
import state from './state'
import form from './form'
import sstate from '../statuses/state'

const Project = {
    oninit(vnode) {
        state.project = { project_users: [], files: [] }
        state.errors = []
        sstate.getAll().then((result) => {
            if (result && result.length > 0) state.project.status_id = result[0].id
        })
    },
    
    view(vnode) {
        let ui = vnode.state
        return m(".projects", [
            m('h1.mb-4', 'New project'),
            form(),
            m('.actions', [
                m('button.btn.btn-primary.mr-2[type=button]', { onclick: state.create }, "Create"),
                m('button.btn.btn-secondary[type=button]', { onclick: () => { m.route.set('/projects') } }, "Cancel")
            ]),
        ])
    }
}

export default Project;