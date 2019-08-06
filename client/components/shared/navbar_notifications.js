import m from 'mithril'
import dropdown from './dropdown'

let state = {
    show: false,
}

const Notifications = {
    view(vnode) {
        /* return m(dropdown, {
            children: [
                m('a.nav-link#navbar-notifications[href=#]', [
                    m('span.fa.fa-bell-o'),
                    m('span.badge.badge-pill.badge-primary', 5)
                ]),
                m('.dropdown-menu', [
                    m('a.dropdown-item[href=#!/statuses]', 'Project Statuses'),
                    m('a.dropdown-item[href=#!/task_steps]', 'Task Steps'),
                    m('a.dropdown-item[href=#!/roles]', 'User Roles'),
                ]),
            ]
        }) */
        return m('li.nav-item.dropdown.mr-2#navbar-notifications', [
            m('a.nav-link.dropdown-toggle[href=#]', [
                m('span.fa.fa-bell-o'),
                m('span.badge.badge-pill.badge-primary', 5)
            ])
        ])
    }
}
export default Notifications;