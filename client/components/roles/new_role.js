import m from 'mithril'
import {
    responseErrors
} from '../../utils/helpers'
import Auth from '../../utils/auth'
import error from '../shared/error'

const state = {
    errors: [],
    name: '',
    onCreateCallback: null,
    onCancelCallback: null,
    setName(name) {
        state.name = name
    },
    dispatch(action, args) {
        state[action].apply(state, args || [])
    },

    create() {
        m.request({
            method: "POST",
            url: "/api/roles",
            body: {
                name: state.name
            },
            headers: {
                Authorization: Auth.authHeader()
            }
        }).then((result) => {
            if (typeof state.onCreateCallback == "function") state.onCreateCallback()
        }).catch((error) => state.errors = responseErrors(error))
    },
    cancel() {
        if (typeof state.onCancelCallback == "function") state.onCancelCallback()
    }
}

const NewRole = {
    oninit(vnode) {
        state.errors = []
        state.name = ''
        state.onCreateCallback = vnode.attrs.onCreate
        state.onCancelCallback = vnode.attrs.onCancel
    },
    onKeyPress(event) {
        if (event.keyCode == 13) state.dispatch("create")
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
                    oninput: function (e) {
                        state.setName(e.target.value)
                    },
                    value: state.name
                }),
                m('.input-group-append', [
                    m('button.btn.btn-outline-success[type=button]', {
                        onclick: () => {
                            state.dispatch("create")
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

export default NewRole;