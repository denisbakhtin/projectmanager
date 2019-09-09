import m from 'mithril'
import {
    responseErrors
} from '../../utils/helpers'
import {
    addSuccess
} from '../shared/notifications'

export default function TaskStep() {
    let step = {},
        steps = [],
        errors = [],

        //requests
        get = () =>
        service.getTaskStep(step.id)
        .then((result) => step = result)
        .catch((error) => errors = responseErrors(error)),

        destroy = () =>
        service.deleteTaskStep(step.id)
        .then((result) => {
            addSuccess("Task step removed.")
            m.route.set('/task_steps', {}, {
                replace: true
            })
        })
        .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            step = {
                id: m.route.param('id')
            }
            get()
        },

        view(vnode) {
            return m(".task_step", [
                step.name ? [
                    m('h1.mb-2', step.name),
                    m('p', 'Is final: ' + step.is_final)
                ] : null,
                m('.actions', [
                    m('button.btn.btn-primary.mr-2[type=button]', {
                        onclick: () => {
                            m.route.set('/task_steps/edit/' + step.id)
                        }
                    }, "Edit"),
                    m('button.btn.btn-secondary.mr-2[type=button]', {
                        onclick: () => {
                            m.route.set('/task_steps')
                        }
                    }, "Back to list"),
                    m('button.btn.btn-outline-danger[type=button]', {
                        onclick: destroy
                    }, "Remove step")
                ])
            ])
        }
    }
}
