import m from 'mithril';
import navbar from './navbar'
import notifications from './notifications'
import side_menu from './side_menu'
import ActiveTask from './active_task'

export default function Layout() {
    return {
        view(vnode) {
            //vnode.attrs -> body component
            return m(".layout.restricted-layout", [
                m("header", m(navbar)),
                m('.content-wrapper', [
                    m(side_menu),
                    m("#main", m(vnode.attrs.child)),
                ]),
                m(notifications),
                m(ActiveTask),
            ]);
        }
    }
}
