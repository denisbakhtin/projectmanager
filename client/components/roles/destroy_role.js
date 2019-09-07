import m from 'mithril'
import {
    responseErrors
} from '../../utils/helpers'
import Auth from '../../utils/auth'
import modal from '../shared/modal'
import error from '../shared/error'
import service from '../../utils/service.js'

export default function DestroyRole() {
    let errors = [],
        id = '',
        name = '',
        onDestroyCallback = null,
        onCancelCallback = null,

        destroy = () =>
        service.deleteRole(id)
        .then((result) => {
            if (typeof onDestroyCallback == "function") onDestroyCallback()
        }).catch((error) => errors = responseErrors(error)),

        cancel = () =>
        (typeof onCancelCallback == "function") ? onCancelCallback() : null

    return {
        oninit(vnode) {
            errors = []
            id = vnode.attrs.role.id
            name = vnode.attrs.role.name
            onDestroyCallback = vnode.attrs.onDestroy
            onCancelCallback = vnode.attrs.onCancel
        },
        onKeyPress(event) {
            if (event.keyCode == 13) destroy()
            if (event.keyCode == 27) cancel()
        },

        view(vnode) {
            return m(modal, {
                title: "Are you sure?",
                body: [
                    m('div', `You are about to permanently remove ${name} role.`),
                    m(error, {
                        errors: errors
                    })
                ],
                buttons: [
                    m('button.btn.btn-primary[type=button]', {
                        onclick: destroy,
                    }, "Remove"),
                    m('button.btn.btn-secondary[type=button]', {
                        onclick: cancel,
                    }, "Cancel")
                ],
                methods: {
                    onEnter: destroy,
                    onEsc: cancel,
                }
            })
        }
    }
}
