import m from 'mithril'

//params: onOk, onCancel - callbacks
//        okText, cancelText - optional button names
//        title - modal title, body - modal body
//        large - modal size
export default function Modal() {
    let onOk,
        onCancel,
        large = false,
        small = false

    return {
        oninit(vnode) {
            onOk = vnode.attrs.onOk ?? (() => null)
            onCancel = vnode.attrs.onCancel ?? (() => null)
            large = vnode.attrs.large ?? false
            small = vnode.attrs.small ?? false
        },
        onKeyPress(event) {
            if (event.keyCode == 13) onOk()
            if (event.keyCode == 27) onCancel()
        },

        view(vnode) {
            return m('.modal.show[tabindex=-1][role=dialog]', {
                oncreate: (el) => el.dom.focus(),
                onkeypress: vnode.attrs.onKeyPress
            }, [
                m('.modal-dialog.modal-dialog-centered[role=document]', {
                    class: (large) ? "modal-lg" : ((small) ? "modal-sm" : "")
                }, [
                    m('.modal-content', [
                        m('.modal-header', [
                            m('h5.modal-title', vnode.attrs.title || "Are you sure?"),
                        ]),
                        m('.modal-body', vnode.attrs.body || ''),
                        m('.modal-footer', [
                            m('button[type=submit].btn.btn-primary', { onclick: onOk }, [
                                m('i.fa.fa-check.mr-1'),
                                vnode.attrs.okText ?? "Submit",
                            ]),
                            m('button[type=button].btn.btn-outline-secondary', { onclick: onCancel }, vnode.attrs.cancelText ?? "Cancel"),
                        ])
                    ])
                ])
            ])
        }
    }
}
