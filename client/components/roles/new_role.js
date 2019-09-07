import m from 'mithril'
import {
    responseErrors
} from '../../utils/helpers'
import Auth from '../../utils/auth'
import error from '../shared/error'
import service from '../../utils/service.js'

export default function NewRole() {
    let errors = [],
        name = '',
        onCreateCallback = null,
        onCancelCallback = null,

        setName = (value) => name = value,

        create = () =>
        service.createRole(name)
        .then((result) => {
            if (typeof onCreateCallback == "function") onCreateCallback()
        }).catch((error) => errors = responseErrors(error)),

        cancel = () =>
        (typeof onCancelCallback == "function") ? onCancelCallback() : null

    return {
        oninit(vnode) {
            errors = []
            name = ''
            onCreateCallback = vnode.attrs.onCreate
            onCancelCallback = vnode.attrs.onCancel
        },
        onKeyPress(event) {
            if (event.keyCode == 13) create()
            if (event.keyCode == 27) cancel()
        },

        view(vnode) {
            return m('.input-group.mb-2', [
                m('.input-group', [
                    m('input.form-control[placeholder="Enter role name"]', {
                        oncreate: (el) => {
                            el.dom.focus()
                        },
                        onkeypress: vnode.attrs.onKeyPress,
                        oninput: (e) => setName(e.target.value),
                        value: name
                    }),
                    m('.input-group-append', [
                        m('button.btn.btn-outline-success[type=button]', {
                            onclick: create,
                        }, m('i.fa.fa-check')),
                        m('button.btn.btn-outline-secondary[type=button]', {
                            onclick: cancel,
                        }, m('i.fa.fa-times'))
                    ])
                ]),
                m(error, {
                    errors: errors
                })
            ])
        }
    }
}
