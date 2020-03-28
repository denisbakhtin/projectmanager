import m from 'mithril'

export default function ButtonMenu() {
    let show = false,

        toggle = () => show = !show

    return {
        oninit: (vnode) => { },
        view: (vnode) => {
            return m('.dropdown', {
                class: show ? 'show' : '',
            }, [
                m('button.btn.btn-default[type=button]', { onclick: toggle }, '...'),
                m('div.dropdown-overlay', { onclick: toggle }),
                vnode.attrs.children
            ])
        }
    }
}
