import m from 'mithril'
import { responseErrors } from '../../utils/helpers'
import Auth from '../../utils/auth'
import modal from '../shared/modal'
import error from '../shared/error'

const state = {
    errors: [],
    id: '',
    name: '',
    onDestroyCallback: null,
    onCancelCallback: null,
    dispatch(action, args) { state[action].apply(state, args || []) },

    destroy() {
        m.request({
            method: "DELETE",
            url: "/api/roles/" + state.id,
            headers: { Authorization: Auth.authHeader() }
        }).then((result) => {
            if (typeof state.onDestroyCallback == "function") state.onDestroyCallback()
        }).catch((error) => state.errors = responseErrors(error))
    },
    cancel() {
        if (typeof state.onCancelCallback == "function") state.onCancelCallback()
    }
}

const DestroyRole = {
    oninit(vnode) {
        state.errors = []
        state.id = vnode.attrs.role.id
        state.name = vnode.attrs.role.name
        state.onDestroyCallback = vnode.attrs.onDestroy
        state.onCancelCallback = vnode.attrs.onCancel
    },
    onKeyPress(event) {
        if (event.keyCode == 13) state.dispatch("destroy")
        if (event.keyCode == 27) state.dispatch("cancel")
    },

    view(vnode) {
        let ui = vnode.state;
        return m(modal, {
            title: "Are you sure?",
            body: [
                m('div', `You are about to permanently remove ${state.name} role.`),
                m(error, { errors: state.errors })
            ],
            buttons: [
                m('button.btn.btn-primary[type=button]', { onclick: () => { state.dispatch("destroy") } }, "Remove"),
                m('button.btn.btn-secondary[type=button]', { onclick: () => { state.dispatch("cancel") } }, "Cancel")
            ],
            methods: {
                onEnter: () => { state.dispatch("destroy") },
                onEsc: () => {state.dispatch("cancel")}
            }
        })
    }
}

export default DestroyRole;