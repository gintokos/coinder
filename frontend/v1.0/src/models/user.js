class User {
    id = 0
    username = ''
    photo_url = ''
    auth_date = ''
    created_at = ''
    
    constructor(data) {
        Object.assign(this, data)
    }
}

export default User