import m from 'mithril'
import error from '../shared/error'
import {
    responseErrors
} from '../../utils/helpers'
import service from '../../utils/service'
import tasks_list from './tasks_list'

const Filters = Object.freeze({
    ALL: (task) => true,
    OPEN: (task) => task.completed == false,
    SOLVED: (task) => task.completed == true,
    EXPIRED: (task) => !!task.end_date && Date.parse(task.end_date) < Date.now(),
});

export default function Tasks() {
    let tasks = [],
        errors = [],

        getAll = () =>
            service.getTasks()
                .then((result) => tasks = result.slice(0))
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            getAll()
        },

        view(vnode) {
            return m(".tasks", [
                m('h1.title', 'Tasks'),
                m(tasks_list, { tasks: tasks, onUpdate: getAll }),
                m(error, { errors: errors }),
                m('.actions.mt-4', [
                    m('button.btn.btn-primary[type=button]', {
                        onclick: () => m.route.set('/tasks/new')
                    }, [
                        m('i.fa.fa-plus.mr-1'),
                        "New task"
                    ])
                ]),
            ])
        }
    }
}
