import m from 'mithril'
import moment from 'moment'

const SideMenu = {
    view(vnode) {
        return m('#side-menu.wide', [
            m('ul.nav.flex-column', [
                m('li.nav-item', m('a.nav-link[href=#!/]', [
                    m('span.fa.fa-home'),
                    m('span', 'Dashboard')
                ])),
                m('li.nav-item', m('a.nav-link[href=#!/projects]', [
                    m('span.fa.fa-list'),
                    m('span', 'Projects')
                ])),
                m('li.nav-item', m('a.nav-link[href=#!/tasks]', [
                    m('span.fa.fa-check-square-o'),
                    m('span', 'Tasks')
                ])),
                m('li.nav-item.dropdown', [
                    m('a.nav-link[href=#]', [
                        m('span.fa.fa-wrench'),
                        m('span', 'Settings')
                    ]),
                    m('.dropdown-menu', [
                        m('a.dropdown-item[href=#!/statuses]', 'Project Statuses'),
                        m('a.dropdown-item[href=#!/task_steps]', 'Task Steps'),
                    ]),
                ]),
                m('li.nav-item', m('a.nav-link[href=#!/statuses]', [
                    m('span.fa.fa-edit'),
                    m('span', 'Project Statuses')
                ])),
                m('li.nav-item', m('a.nav-link[href=#!/task_steps]', [
                    m('span.fa.fa-cubes'),
                    m('span', 'Task Steps')
                ])),
                m('li.nav-item', m('a.nav-link[href=#!/users]', [
                    m('span.fa.fa-users'),
                    m('span', 'Users')
                ])),
                m('li.nav-item', m('a.nav-link[href=#!/roles]', [
                    m('span.fa.fa-user-o'),
                    m('span', 'User Roles')
                ])),
                m('li.nav-item', m('a.nav-link[href=#!/manage]', [
                    m('span.fa.fa-key'),
                    m('span', 'Manage Account')
                ])),
                m('li.nav-item', m('a.nav-link[href=#!/logout]', [
                    m('span.fa.fa-sign-out'),
                    m('span', [
                        m('span.fa.fa-home'),
                        m('span', 'Logout')
                    ])
                ])),
            ]),
            m('.bottom-links.row', [
                m('a.d-inline-block.col[href=#]', m('span.fa.fa-lock')),
                m('a.d-inline-block.col[href=#]', m('span.fa.fa-download')),
                m('a.d-inline-block.col[href=#]', m('span.fa.fa-globe')),
                m('a.d-inline-block.col[href=#]', m('span.fa.fa-phone')),
            ])
        ])
    }
}

export default SideMenu;