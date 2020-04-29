import m from 'mithril'
import error from '../shared/error'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors
} from '../../utils/helpers'
import service from '../../utils/service.js'

export default function Setting() {
    let errors = [],
        setting = {},
        isNew = true,

        setCode = (code) => setting.code = code,
        setTitle = (title) => setting.title = title,
        setValue = (value) => setting.value = value,

        get = (id) =>
            service.getSetting(id)
                .then((result) => setting = result)
                .catch((error) => errors = responseErrors(error)),

        create = () =>
            service.createSetting(setting)
                .then((result) => {
                    addSuccess("Setting created.")
                    m.route.set('/settings')
                })
                .catch((error) => errors = responseErrors(error)),

        update = () =>
            service.updateSetting(setting.id, setting)
                .then((result) => {
                    addSuccess("Setting updated.")
                    m.route.set('/settings')
                })
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            isNew = (m.route.param('id') == undefined)
            if (!isNew)
                get(m.route.param('id'))
        },

        view(vnode) {
            return m(".settings", [
                m('h1.title.mb-4', (isNew) ? 'New setting' : 'Edit setting'),
                m('.form-group.form-row', [
                    m('.col', [
                        m('label', 'Setting code'),
                        m('input.form-control[type=text]', {
                            oncreate: (el) => el.dom.focus(),
                            oninput: (e) => setCode(e.target.value),
                            value: setting.code
                        })
                    ]),
                    m('.col', [
                        m('label', 'Title'),
                        m('input.form-control[type=text]', {
                            oninput: (e) => setTitle(e.target.value),
                            value: setting.title
                        })
                    ])
                ]),
                m('.form-group', [
                    m('label', "Value"),
                    m('textarea.form-control', {
                        oninput: (e) => setValue(e.target.value),
                        value: setting.value
                    })
                ]),
                m('.mb-2', m(error, { errors: errors })),
                m('.actions', [
                    m('button.btn.btn-primary.mr-2[type=button]', {
                        onclick: (isNew) ? create : update
                    }, "Save"),
                    m('button.btn.btn-secondary[type=button]', {
                        onclick: () => window.history.back()
                    }, "Cancel")
                ]),
            ])
        }
    }
}
