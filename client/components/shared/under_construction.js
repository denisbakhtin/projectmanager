import m from 'mithril'

export default function UnderConstruction() {
    return {
        oninit(vnode) { },

        view(vnode) {
            return m(".under-construction", [
                m('img[src=/public/images/sloth.png]'),
                m('h1', 'This page is under construction')
            ])
        }
    }
}
