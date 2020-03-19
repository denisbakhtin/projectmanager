import m from 'mithril'
import error from '../shared/error'
import {
    responseErrors
} from '../../utils/helpers'
import service from '../../utils/service'
import tasks_item from './tasks_item.js'

const Filters = Object.freeze({
    ALL: (task) => true,
    OPEN: (task) => task.completed == false,
    SOLVED: (task) => task.completed == true
});

export default function Tasks() {
    let tasks = [],
        errors = [],
        filter,

        activeClass = (fil) => (filter === Filters[fil]) ? "active" : "",
        setFilter = (fil) => {
            localStorage.taskFilter = fil
            filter = Filters[fil]
        },
        getFilter = () => localStorage.taskFilter,

        getAll = () =>
            service.getTasks()
                .then((result) => tasks = result.slice(0))
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            getAll()
            filter = Filters[getFilter() ?? "ALL"] ?? Filters["ALL"]
        },

        view(vnode) {
            let filteredTasks = tasks.filter(filter)

            return m(".tasks", [
                m('h1.title', 'Tasks'),
                m('.filters', [
                    m('button.btn.btn-link', { class: activeClass("ALL"), onclick: () => setFilter("ALL") }, "All"),
                    m('button.btn.btn-link', { class: activeClass("OPEN"), onclick: () => setFilter("OPEN") }, "Open"),
                    m('button.btn.btn-link', { class: activeClass("SOLVED"), onclick: () => setFilter("SOLVED") }, "Solved"),
                ]),
                (filteredTasks && filteredTasks.length > 0) ? m('ul.dashboard-box.box-list',
                    filteredTasks.map((task) => m(tasks_item, { key: task.id, task: task, onUpdate: getAll }))
                ) : m('p.text-muted', 'The list is empty'),
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
