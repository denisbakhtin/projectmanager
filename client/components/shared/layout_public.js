import m from 'mithril';
import navbar from './navbar_public'
import notifications from './notifications'

const Layout = {
    //vnode.attrs -> body component
    view(vnode) {
        return m(".layout.public-layout", [
            m("header", m(navbar)),
            m("#main.container", m(vnode.attrs)),
            m(notifications),
        ]);
    }
}
export default Layout;