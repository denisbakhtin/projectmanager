import m from 'mithril';
import navbar from './navbar'
import notifications from './notifications'
import side_menu from './side_menu'

const Layout = {
    //vnode.attrs -> body component
    view(vnode) {
        return m(".layout.restricted-layout", [
            m("header", m(navbar)),
            m(side_menu),
            m("#main", m(vnode.attrs)),
            m(notifications),
        ]);
    }
}
export default Layout;