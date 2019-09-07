import m from 'mithril'

export default function Modal() {
    let onEnter = undefined,
        onEsc = undefined
    return {
        oninit(vnode) {
            if (vnode.attrs.methods) {
                onEnter = vnode.attrs.methods.onEnter
                onEsc = vnode.attrs.methods.onEsc
            }
        },
        onKeyPress(event) {
            if (event.keyCode == 13)
                if (typeof onEnter == 'function') onEnter()
            if (event.keyCode == 27)
                if (typeof onEsc == 'function') onEsc()
        },

        view(vnode) {
            let ui = vnode.state;
            return m('modal[tabindex=-1][role=dialog]', {
                oncreate: (el) => {
                    el.dom.focus()
                },
                onkeypress: ui.onKeyPress
            }, [
                m('.modal-dialog[role=document]', [
                    m('.modal-content', [
                        m('.modal-header', [
                            m('h5.modal-title', vnode.attrs.title || "Are you sure?"),
                        ]),
                        m('.modal-body', vnode.attrs.body || ''),
                        m('.modal-footer', vnode.attrs.buttons)
                    ])
                ])
            ])
        }
    }
}
