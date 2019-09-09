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
    if (error.statusCode == 401) return ["Authorization required."]
    if (error.statusCode == 404) return ["Sorry, the request url you tried cannot be found."]
    if (!!error.errors) error.errors.slice(0)
    if (!!error.error) return [error.error]
    if (!!error.message) return [error.message]
    return ["Your request resulted in an error."]
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
