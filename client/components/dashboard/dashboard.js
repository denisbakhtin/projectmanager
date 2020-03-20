import m from 'mithril'
import { ProjectsCountWidget } from '../projects/projects_widgets'
import { TasksCountWidget, LatestTasksWidget, LatestTaskLogsWidget } from '../tasks/tasks_widgets'
import { CategoriesCountWidget } from '../categories/categories_widgets'
import { SessionsCountWidget } from '../sessions/sessions_widgets'
import { UsersCountWidget } from '../users/users_widgets'
import auth from '../../utils/auth'

export default function Dashboard() {
    return {
        oninit(vnode) { },

        view(vnode) {
            return m(".dashboard", [
                m('h1.title', 'Dashboard'),
                m('.row.count-widgets', [
                    m('.col-sm-4.col-md-3.col-lg-2.mb-3', m(ProjectsCountWidget)),
                    m('.col-sm-4.col-md-3.col-lg-2.mb-3', m(TasksCountWidget)),
                    m('.col-sm-4.col-md-3.col-lg-2.mb-3', m(CategoriesCountWidget)),
                    m('.col-sm-4.col-md-3.col-lg-2.mb-3', m(SessionsCountWidget)),
                    auth.isAdmin() ? m('.col-sm-4.col-md-3.col-lg-2.mb-3', m(UsersCountWidget)) : null,
                ]),
                m('.row.task-widgets', [
                    m('.col-sm-6.mb-3', m(LatestTasksWidget)),
                    m('.col-sm-6.mb-3', m(LatestTaskLogsWidget)),
                ]),
            ])
        }
    }
}
