import m from 'mithril'
import dropdown from './dropdown'
import state from './state'
import Auth from '../../utils/auth'

export default function SideMenu() {
    return {
        oninit(vnode) {
            state.getFavoriteProjects()
        },

        view(vnode) {
            return m('#side-menu', {
                class: state.sidebarCollapsed ? 'narrow' : 'wide'
            }, [
                m('ul.nav.flex-column', [
                    m('li.nav-item', m('a.nav-link[href=#!/]', [
                        m('span.fa.fa-home'),
                        m('span.title', 'Dashboard')
                    ])),
                    m('li.nav-item', m('a.nav-link[href=#!/projects]', [
                        m('span.fa.fa-list'),
                        m('span.title', 'Projects')
                    ])),
                    (state.favoriteProjects && state.favoriteProjects.length > 0) ?
                        state.favoriteProjects.map((project) => m('li.nav-item',
                            m('a.nav-link.favorite-project', { href: "#!/projects/" + project.id }, [
                                m('span.fa.fa-star'),
                                m('span.title', project.name)
                            ]))
                        ) : null,
                    m('li.nav-item', m('a.nav-link[href=#!/tasks]', [
                        m('span.fa.fa-check-square-o'),
                        m('span.title', 'Tasks')
                    ])),
                    m('li.nav-item', m('a.nav-link[href=#!/reports/spent]', [
                        m('span.fa.fa-bar-chart'),
                        m('span.title', 'Spent Report')
                    ])),
                    m('li.nav-item', m('a.nav-link[href=#!/sessions]', [
                        m('span.fa.fa-check'),
                        m('span.title', 'Sessions')
                    ])),

                    m(dropdown, {
                        id: 'settings-dropdown',
                        keepState: true,
                        children: [
                            m('a.nav-link[href=#]', [
                                m('span.fa.fa-wrench'),
                                m('span.title', 'Settings')
                            ]),
                            m('.dropdown-menu', [
                                m('a.dropdown-item[href=#!/categories]', 'Categories'),
                                //m('a.dropdown-item[href=#!/roles]', 'User Roles'),
                                (Auth.isAdmin()) ? m('a.dropdown-item[href=#!/pages]', 'Pages') : null,
                                (Auth.isAdmin()) ? m('a.dropdown-item[href=#!/settings]', 'Site Settings') : null,
                                (Auth.isAdmin()) ? m('a.dropdown-item[href=#!/users]', 'Users') : null,
                                m('a.dropdown-item[href=#!/manage]', 'Manage Account'),
                            ]),
                        ]
                    }),
                    m('li.nav-item', m('a.nav-link[href=#!/support]', [
                        m('span.fa.fa-question-circle'),
                        m('span.title', 'Contact Support')
                    ])),
                    m('li.nav-item', m('a.nav-link[href=#!/logout]', [
                        m('span.fa.fa-sign-out'),
                        m('span.title', 'Logout')
                    ])),
                ]),
                m('.bottom-links.row', [
                    m('a.d-inline-block.col[href=#!/lock]', m('span.fa.fa-lock')),
                    m('a.d-inline-block.col[href=#!/download]', m('span.fa.fa-download')),
                    m('a.d-inline-block.col[href=#!/web_site]', m('span.fa.fa-globe')),
                    m('a.d-inline-block.col[href=#!/phone]', m('span.fa.fa-phone')),
                ])
            ])
        }
    }
}
