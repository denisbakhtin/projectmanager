import m from 'mithril';
import navbar from './navbar_public'
import notifications from './notifications'

export default function Layout() {
    return {
        //vnode.attrs -> body component
        view(vnode) {
            return m(".layout.public-layout", [
                m("header", m(navbar)),
                m("#main.container", m(vnode.attrs.child)),
                m(notifications),
            ]);
        }
    }
}
