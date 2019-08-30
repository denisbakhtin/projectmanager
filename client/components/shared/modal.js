import m from 'mithril'

const Modal = {
    onEnter: undefined,
    onEsc: undefined,
    oninit(vnode) {
        if (vnode.attrs.methods) Modal.onEnter = vnode.attrs.methods.onEnter
        if (vnode.attrs.methods) Modal.onEsc = vnode.attrs.methods.onEsc
    },
    onKeyPress(event) {
        if (event.keyCode == 13)
            if (typeof Modal.onEnter == 'function') Modal.onEnter()
        if (event.keyCode == 27)
            if (typeof Modal.onEsc == 'function') Modal.onEsc()
    },

    view(vnode) {
        let ui = vnode.state;
        return m('modal[tabindex=-1][role=dialog]', {
            oncreate: (el) => { el.dom.focus() },
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

export default Modal;