import m from 'mithril'
import error from '../shared/error'
import service from '../../utils/service'
import {
    responseErrors
} from '../../utils/helpers'

export default function Settings() {
    let errors = [],
        settings = [],

        getAll = () =>
            service.getSettings()
                .then((result) => settings = result.slice(0))
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            getAll()
        },

        view(vnode) {
            return m(".settings", [
                m('h1.mb-4', 'Site settings'),
                m('table.table', [
                    m('thead', [
                        m('tr', [
                            m('th[scope=col]', 'Code'),
                            m('th[scope=col]', 'Title'),
                            m('th[scope=col]', 'Value'),
                            m('th.shrink.text-center[scope=col]', 'Actions')
                        ])
                    ]),
                    m('tbody', [
                        settings ?
                            settings.map((setting) => {
                                return m('tr', {
                                    key: setting.id
                                }, [
                                    m('td', setting.code),
                                    m('td', setting.title),
                                    m('td', setting.value),
                                    m('td.shrink.text-center', [

                                        m('button.btn.btn-outline-primary.btn-sm.mr-2[type=button]', {
                                            onclick: () => m.route.set('/settings/edit/' + setting.id)
                                        }, m('i.fa.fa-pencil')),
                                        m('button.btn.btn-outline-danger.btn-sm[type=button]', {
                                            onclick: () => service.deleteSetting(setting.id)
                                        }, m('i.fa.fa-times')),
                                    ])
                                ])
                            }) : null
                    ])
                ]),
                m(error, { errors: errors }),
                m('.actions.mt-4', [
                    m('button.btn.btn-primary[type=button]', {
                        onclick: () => m.route.set('/settings/new')
                    }, "New setting")
                ]),
            ])
        }
    }
}
