import m from 'mithril'
import error from '../shared/error'
import {
    addSuccess
} from '../shared/notifications'
import state from './state'
import steps_state from '../tasksteps/state'

export default function form() {
    //TODO: MOVE TO ONINIT ASAP, it forces infinite redraw loops
    steps_state.getAll()
    return [
        m('.form-group', [
            m('label', 'Task name'),
            m('input.form-control[type=text]', {
                oncreate: (el) => {
                    el.dom.focus()
                },
                oninput: (e) => state.setName(e.target.value),
                value: state.task.name
            })
        ]),
        m('.form-group', [
            m('label', 'State'),
            m('select.form-control', {
                    onchange: (e) => state.setStepID(e.target.value),
                    value: state.task.task_step_id
                }, steps_state.steps ?
                steps_state.steps.map((step) => {
                    return m('option', {
                        value: step.id
                    }, step.name)
                }) : null)
        ]),
        m('.form-group', [
            m('label', "Description"),
            m('textarea.form-control', {
                oninput: (e) => state.setDescription(e.target.value),
                value: state.task.description
            })
        ]),
        m('.form-group', [
            m('label', 'Assigned user'),
            m('select.form-control', {
                    onchange: (e) => state.setProjectUserID(e.target.value),
                    value: state.task.project_user_id
                }, steps_state.steps ?
                steps_state.steps.map((step) => {
                    return m('option', {
                        value: step.id
                    }, step.name)
                }) : null)
        ]),
        m('.form-group', [
            m('label', 'Task name'),
            m('input.form-control[type=text]', {
                oninput: (e) => state.setName(e.target.value),
                value: state.task.name
            })
        ]),
        m('.mb-2', m(error, {
            errors: state.errors
        })),
    ]
}