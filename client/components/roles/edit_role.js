import m from 'mithril'
import {
    responseErrors
} from '../../utils/helpers'
import Auth from '../../utils/auth'
import error from '../shared/error'

const state = {
    errors: [],
    id: '',
    name: '',
    onUpdateCallback: null,
    onCancelCallback: null,
    setName(name) {
        state.name = name
    },
    dispatch(action, args) {
        state[action].apply(state, args || [])
    },

    update() {
        m.request({
            method: "PUT",
            url: "/api/roles/" + state.id,
            body: {
                id: state.id,
                name: state.name
            },
            headers: {
                Authorization: Auth.authHeader()
            }
        }).then((result) => {
            if (typeof state.onUpdateCallback == "function") state.onUpdateCallback()
        }).catch((error) => state.errors = responseErrors(error))
    },
    cancel() {
        if (typeof state.onCancelCallback == "function") state.onCancelCallback()
    }
}

const EditRole = {
    oninit(vnode) {
        state.errors = []
        state.id = vnode.attrs.role.id
        state.name = vnode.attrs.role.name
        state.onUpdateCallback = vnode.attrs.onUpdate
        state.onCancelCallback = vnode.attrs.onCancel
    },
    onKeyPress(event) {
        if (event.keyCode == 13) state.dispatch("update")
        if (event.keyCode == 27) state.dispatch("cancel")
    },

    view(vnode) {
        let ui = vnode.state;
        return m('.input-group.mb-2', [
            m('.input-group', [
                m('input.form-control[placeholder="Enter role name"]', {
                    oncreate: (el) => {
                        el.dom.focus()
                    },
                    onkeypress: ui.onKeyPress,
                    oninput: (e) => state.setName(e.target.value),
                    value: state.name
                }),
                m('.input-group-append', [
                    m('button.btn.btn-outline-success[type=button]', {
                        onclick: () => {
                            state.dispatch("update")
                        }
                    }, m('i.fa.fa-check')),
                    m('button.btn.btn-outline-secondary[type=button]', {
                        onclick: () => {
                            state.dispatch("cancel")
                        }
                    }, m('i.fa.fa-times'))
                ])
            ]),
            m(error, {
                errors: state.errors
            })
        ])
    }
}

export default EditRole;