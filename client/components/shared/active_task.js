import m from 'mithril'
import {
    responseErrors,
    humanSpent
} from '../../utils/helpers'
import service from '../../utils/service'
import { addDanger } from '../shared/notifications'
import global_state from '../shared/state'

let state = {
    task_log: undefined,
    timerId: undefined,
    startedWhen: undefined,
    onStop: undefined,
    stop() {
        state.update().then((result) => {
            state.task_log = undefined
            state.startedWhen = undefined
            clearInterval(state.timerId)
            state.onStop()
        })
    },
    start: () =>
        service.createTaskLog(state.task_log)
            .then((result) => {
                state.task_log.id = result.id //need only id 
                state.startedWhen = Date.now()
                state.timerId = setInterval(state.update, 60000);
            }).catch((error) => addDanger(responseErrors(error).join(', ')))
    ,
    update: () => {
        state.task_log.minutes = Math.ceil((Date.now() - state.startedWhen) / 60000)
        return service.updateTaskLog(state.task_log.id, state.task_log)
            .catch((error) => addDanger(responseErrors(error).join(', ')))
    },
    spent: () => humanSpent(state.task_log.minutes)
}

export function startTask(task, onStop) {
    state.task_log = { task_id: task.id, minutes: 0, task: task }
    state.onStop = (typeof onStop === 'function') ? onStop : (() => null)
    state.start()
}

export default function ActiveTask() {
    return {
        oninit(vnode) { },
        view(vnode) {
            return state.task_log ?
                m('.active-task', { class: global_state.sidebarCollapsed ? "wide" : "narrow" }, [
                    m('a', { href: '#!/tasks/' + state.task_log.task.id }, state.task_log.task.name),
                    (state.spent() != '') ? m('span.text-muted.ml-2', state.spent()) : null,
                    m('button.btn.btn-sm.btn-secondary.ml-3', { onclick: state.stop }, [
                        m('i.fa.fa-stop.mr-1'),
                        m('span', 'Stop'),
                    ])
                ]) : null
        }
    }
}
