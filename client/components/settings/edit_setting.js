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
    var errors = [],
        setting = {},
        isNew = true

    function setCode(code) {
        setting.code = code
    }
    function setTitle(title) {
        setting.title = title
    }
    function setValue(value) {
        setting.value = value
    }
    function validate() {
        errors = []
        if (!setting.name)
            errors.push("Setting code is required.")
        return errors.length == 0
    }
    function get() {
        return service.getSetting(setting.id)
            .then((result) => setting = result)
            .catch((error) => errors = responseErrors(error))
    }
    function create() {
        return service.createSetting(setting)
            .then((result) => {
                addSuccess("Setting created.")
                m.route.set('/settings')
            })
            .catch((error) => errors = responseErrors(error))
    }
    function update() {
        return service.updateSetting(setting.id, setting)
            .then((result) => {
                addSuccess("Setting updated.")
                m.route.set('/settings')
            })
            .catch((error) => errors = responseErrors(error))
    }

    return {
        oninit(vnode) {
            if (m.route.param('id')) {
                isNew = false
                setting = { id: m.route.param('id')}
                get()
            } else
                setting = {}
            errors = []
        },

        view(vnode) {
            return m(".settings", [
                m('h1.mb-4', (isNew) ? 'New setting' : 'Edit setting'),
                m('.form-group.form-row', [
                    m('.col', [
                        m('label', 'Setting code'),
                        m('input.form-control[type=text]', {
                            oncreate: (el) => {
                                el.dom.focus()
                            },
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
                m('.mb-2', m(error, {
                    errors: errors
                })),
                m('.actions', [
                    m('button.btn.btn-primary.mr-2[type=button]', { onclick: (isNew) ? create : update }, "Save"),
                    m('button.btn.btn-secondary[type=button]', { onclick: () => { window.history.back() } }, "Cancel")
                ]),
            ])
        }
    }
}
