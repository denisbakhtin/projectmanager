import m from 'mithril'
import {
    responseErrors
} from '../../utils/helpers'
import Auth from '../../utils/auth'
import error from '../shared/error'
import service from '../../utils/service.js'

export default function EditRole() {
    let errors = [],
        id = '',
        name = '',
        onUpdateCallback = null,
        onCancelCallback = null,

        setName = (value) => name = value,

        update = () =>
        service.updateRole(id, name)
        .then((result) => {
            if (typeof onUpdateCallback == "function") onUpdateCallback()
        }).catch((error) => errors = responseErrors(error)),

        cancel = () =>
        (typeof onCancelCallback == "function") ? onCancelCallback() : null

    return {
        oninit(vnode) {
            errors = []
            id = vnode.attrs.role.id
            name = vnode.attrs.role.name
            onUpdateCallback = vnode.attrs.onUpdate
            onCancelCallback = vnode.attrs.onCancel
        },
        onKeyPress(event) {
            if (event.keyCode == 13) update()
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
                            onclick: update,
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
