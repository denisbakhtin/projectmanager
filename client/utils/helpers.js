export function emailIsValid(email) {
    var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(email);
}

export function guid() {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, (c) => {
        var r = Math.random() * 16 | 0,
            v = c == 'x' ? r : (r & 0x3 | 0x8);
        return v.toString(16);
    });
}

export function responseErrors(error) {
    if (!error) return []
    if (!!error.message) return [error.message]
    if (error.code == 401) return ["Authorization required with appropriate privileges."]
    if (error.code == 404) return ["Sorry, the request url you tried cannot be found."]

    return ["Your request resulted in an error."]
}

export function isZeroDate(date) {
    return String(date).startsWith("0001-01-01")
}

export function humanDate(date) {
    if (typeof date === 'string') date = new Date(date)
    const ye = new Intl.DateTimeFormat('en', { year: 'numeric' }).format(date)
    const mo = new Intl.DateTimeFormat('en', { month: 'short' }).format(date)
    const da = new Intl.DateTimeFormat('en', { day: '2-digit' }).format(date)

    return `${da} ${mo}, ${ye}`
}

export function humanSpent(minutes, long) {
    if (!minutes || minutes == 0) return ''
    let m = minutes % 60
    let h = Math.floor(minutes / 60)
    let res
    if (long)
        res = ((h > 0) ? `${h} hours ` : '') + ((m > 0) ? `${m} minutes` : '')
    else
        res = ((h > 0) ? `${h}h ` : '') + ((m > 0) ? `${m}m` : '')

    return res.trim()
}

export function humanTaskSpent(task, long) {
    return humanSpent(taskSpent(task), long)
}

function taskSpent(task) {
    let logs = task.task_logs ?? []
    return logs.reduce((total, cur) => total + cur.minutes, 0)
}

export function humanProjectSpent(project, long) {
    return humanSpent(projectSpent(project), long)
}

function projectSpent(project) {
    let tasks = project.tasks ?? []
    return tasks.reduce((total, cur) => total + taskSpent(cur), 0)
}

export function humanAllProjectsSpent(projects, long) {
    projects = projects ?? []
    let spent = projects.reduce((total, cur) => total + projectSpent(cur), 0)
    return humanSpent(spent, long)
}

export function humanSessionSpent(session, long) {
    let logs = session.task_logs ?? []
    let spent = logs.reduce((total, cur) => total + cur.minutes, 0)
    return humanSpent(spent, long)
}

export function ISODateToHtml5(datestr, blank) {
    if (datestr) {
        let date = new Date(datestr)
        return `${date.getFullYear()}-${zeroLeadingMonth(date)}-${zeroLeadingDay(date)}`
    }
    return blank
}

function zeroLeadingMonth(date) {
    let month = "0" + (date.getMonth() + 1)
    return month.slice(-2)
}

function zeroLeadingDay(date) {
    let day = "0" + date.getDate()
    return day.slice(-2)
}

export function entityUrl(entity, id) {
    switch (entity) {
        case 'project':
            return '#!/projects' + (id) ? '/' + id : '';
        case 'task':
            return '#!/tasks' + (id) ? '/' + id : '';
        default:
            return '#!/';
    }
}

export function groupLogsByProject(logs) {
    return groupTasksByProjects(groupLogsByTask(logs))
}

function groupLogsByTask(logs) {
    if (!Array.isArray(logs)) return []

    let tasks = {}
    logs.forEach((log) => {
        log.task.task_logs = []
        let task = tasks[log.task.id] ?? log.task
        task.task_logs.push(log)
        tasks[task.id] = task
    })
    return Object.values(tasks)
}

function groupTasksByProjects(tasks) {
    if (!Array.isArray(tasks)) return []

    let projects = {}
    tasks.forEach((task) => {
        task.project.tasks = []
        let project = projects[task.project.id] ?? task.project
        project.tasks.push(task)
        projects[project.id] = project
    })
    return Object.values(projects)
}

export function firstOfWeek() {
    let date = new Date(Date.now())
    let day = date.getDay()
    let diff = date.getDate() - day + (day == 0 ? -6 : 1); // adjust when day is sunday
    return new Date(date.setDate(diff));
}

export function firstOfMonth() {
    let date = new Date(Date.now())
    return new Date(date.getFullYear(), date.getMonth(), 1)
}

export function minusDays(days) {
    let date = new Date(Date.now())
    return new Date(date.setDate(date.getDate() - days))
}