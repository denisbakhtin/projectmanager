import m from 'mithril'
import error from '../shared/error'
import {
    addSuccess
} from '../shared/notifications'
import state from './state'

export default function form() {
    return [
        m('.form-group', [
            m('label', 'Step name'),
            m('input.form-control[type=text]', {
                oncreate: (el) => {
                    el.dom.focus()
                },
                oninput: (e) => state.setName(e.target.value),
                value: state.step.name
            })
        ]),
        m('.form-check', [
            m('input#isfinal.form-check-input[type=checkbox]', {
                oninput: (e) => state.setIsfinal(e.target.value),
                checked: state.step.is_final
            }),
            m('label.form-check-label[for=isfinal]', 'Is final')
        ]),
        m('.form-group w-25', [
            m('label', 'Order'),
            m('input.form-control[type=number][min=0]', {
                oninput: (e) => state.setOrder(e.target.value),
                value: state.step.order
            })
        ]),
        m('.mb-2', m(error, {
            errors: state.errors
        })),
    ]
}