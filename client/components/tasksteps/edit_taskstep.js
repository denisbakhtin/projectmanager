import m from 'mithril'
import error from '../shared/error'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors
} from '../../utils/helpers'

export default function TaskStep() {
    let step = {},
        steps = [],
        errors = [],
        isNew = true,
        setName = (name) => step.name = name,
        setIsfinal = (is_final) => step.is_final = is_final,
        setOrd = (ord) => step.ord = ord,
        validate = () => {
            errors = []
            if (!step.name)
                errors.push("Step name is required.")
            return errors.length == 0
        },

        //requests
        get = () =>
        service.getTaskStep(step.id)
        .then((result) => step = result)
        .catch((error) => errors = responseErrors(error)),

        create = () =>
        service.createTaskStep(step)
        .then((result) => {
            addSuccess("Task step created.")
            m.route.set('/task_steps')
        })
        .catch((error) => errors = responseErrors(error)),

        update = () =>
        service.updateTaskStep(step.id, step)
        .then((result) => {
            addSuccess("Task step updated.")
            m.route.set('/task_steps')
        })
        .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            if (m.route.param('id')) {
                isNew = false
                step = {
                    id: m.route.param('id')
                }
                get()
            }
        },

        view(vnode) {
            return m(".task_step", [
                m('h1.mb-4', (isNew) ? "Create task step" : 'Edit task step'),
                m('.form-group', [
                    m('label', 'Step name'),
                    m('input.form-control[type=text]', {
                        oncreate: (el) => {
                            el.dom.focus()
                        },
                        oninput: (e) => setName(e.target.value),
                        value: step.name
                    })
                ]),
                m('.form-check', [
                    m('input#isfinal.form-check-input[type=checkbox]', {
                        oninput: (e) => setIsfinal(e.target.value),
                        checked: step.is_final
                    }),
                    m('label.form-check-label[for=isfinal]', 'Is final')
                ]),
                m('.form-group w-25', [
                    m('label', 'Order'),
                    m('input.form-control[type=number][min=0]', {
                        oninput: (e) => setOrd(e.target.value),
                        value: step.ord
                    })
                ]),
                m('.mb-2', m(error, {
                    errors: errors
                })),
                m('.actions', [
                    m('button.btn.btn-primary.mr-2[type=button]', {
                        onclick: (isNew) ? create : update
                    }, "Save"),
                    m('button.btn.btn-secondary[type=button]', {
                        onclick: () => {
                            window.history.back()
                        }
                    }, "Cancel")
                ]),
            ])
        }
    }
}
