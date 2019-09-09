import './css/site.scss'
import m from 'mithril'
import 'moment'
import users from './components/users/users'
import login from './components/account/login'
import logout from './components/account/logout'
import roles from './components/roles/roles'
import register from './components/account/register'
import activationNotice from './components/account/activation_notice'
import activate from './components/account/activate'
import reset from './components/account/reset'
import resetNotice from './components/account/reset_notice'
import manage from './components/account/manage'
import home from './components/home'

import projects from './components/projects/projects'
import edit_project from './components/projects/edit_project'
import project from './components/projects/project'

import statuses from './components/statuses/statuses'
import edit_status from './components/statuses/edit_status'
import status from './components/statuses/status'

import settings from './components/settings/settings'
import edit_setting from './components/settings/edit_setting.js'

import tasksteps from './components/tasksteps/tasksteps'
import edit_taskstep from './components/tasksteps/edit_taskstep'
import taskstep from './components/tasksteps/taskstep'

import tasks from './components/tasks/tasks'
import edit_task from './components/tasks/edit_task'
// import task from './components/tasks/task'

import layout from './components/shared/layout'
import public_layout from './components/shared/layout_public'
import Auth from './utils/auth'

const app_root = document.getElementById("app-root");

//render component with layout
function withLayout(comp) {
    return {
        view: () => Auth.isLoggedIn() ? m(layout, {
            child: comp
        }) : m(public_layout, {
            child: comp
        })
    }
}

//Authorization route filter
function checkAuthorized(comp) {
    return function(args, path) {
        if (!Auth.isLoggedIn()) {
            localStorage.returnURL = path
            m.route.set("/login")
        } else return withLayout(comp)
    }
}

//route wrapper
function route(comp, requiresAuth = true) {
    return (requiresAuth) ? {
        onmatch: checkAuthorized(comp)
    } : withLayout(comp)
}

const routes = {
    '/': route(home, false),
    '/login': route(login, false),
    '/register': route(register, false),
    '/activation_notice': route(activationNotice, false),
    '/activate/:token': route(activate, false),
    '/reset': route(reset, false),
    '/reset/:token': route(reset, false),
    '/reset_notice': route(resetNotice, false),
    '/manage': route(manage),
    '/roles': route(roles),
    '/users': route(users),

    '/projects': route(projects),
    '/projects/new': route(edit_project),
    '/projects/edit/:id': route(edit_project),
    '/projects/:id': route(project),

    '/statuses': route(statuses),
    '/statuses/new': route(edit_status),
    '/statuses/edit/:id': route(edit_status),
    '/statuses/:id': route(status),

    '/task_steps': route(tasksteps),
    '/task_steps/new': route(edit_taskstep),
    '/task_steps/edit/:id': route(edit_taskstep),
    '/task_steps/:id': route(taskstep),

    '/tasks': route(tasks),
    '/tasks/new': route(edit_task),
    '/tasks/edit/:id': route(edit_task),
    // '/tasks/:id': route(task),

    '/settings': route(settings),
    '/settings/new': route(edit_setting),
    '/settings/edit/:id': route(edit_setting),

    '/logout': route(logout, false),
};



m.route(app_root, "/", routes)

//turn on webpack hot reload
if (module.hot) {
    module.hot.accept()
}
