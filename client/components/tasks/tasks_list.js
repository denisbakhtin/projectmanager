import m from 'mithril'
import tasks_item from './tasks_item.js'

const Filters = Object.freeze({
    ALL: (task) => true,
    OPEN: (task) => task.completed == false,
    SOLVED: (task) => task.completed == true,
    EXPIRED: (task) => !!task.end_date && Date.parse(task.end_date) < Date.now(),
});

export default function TasksList() {
    let onUpdate,
        filter,

        activeClass = (fil) => (filter === Filters[fil]) ? "active" : "",
        setFilter = (fil) => {
            localStorage.taskFilter = fil
            filter = Filters[fil]
        },
        getFilter = () => localStorage.taskFilter

    return {
        oninit(vnode) {
            onUpdate = vnode.attrs.onUpdate ?? (() => null)
            filter = Filters[getFilter() ?? "ALL"] ?? Filters["ALL"]
        },

        view(vnode) {
            let tasks = vnode.attrs.tasks ?? []
            let filteredTasks = tasks.filter(filter)

            return [
                m('.filters', [
                    m('button.btn.btn-link', { class: activeClass("ALL"), onclick: () => setFilter("ALL") }, "All"),
                    m('button.btn.btn-link', { class: activeClass("OPEN"), onclick: () => setFilter("OPEN") }, "Open"),
                    m('button.btn.btn-link', { class: activeClass("SOLVED"), onclick: () => setFilter("SOLVED") }, "Solved"),
                    m('button.btn.btn-link', { class: activeClass("EXPIRED"), onclick: () => setFilter("EXPIRED") }, "Expired"),
                ]),
                (filteredTasks && filteredTasks.length > 0) ? m('ul.dashboard-box.box-list',
                    filteredTasks.map((task) => m(tasks_item, { key: task.id, task: task, onUpdate: onUpdate, onOpenClick: () => setFilter("OPEN"), onExpiredClick: () => setFilter("EXPIRED") }))
                ) : m('p.text-muted', 'The list is empty'),
            ]
        }
    }
}
