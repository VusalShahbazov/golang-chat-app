
export default {
    data() {
        return {

            pagination :{
                page : 1,
                per_page : 12,
                total : 0,
                loading : false
            }
        }
    },
    computed: {
        pagination_length(){
            return Math.ceil(this.pagination.total/this.pagination.per_page)
        },
        nex_page (){
            return this.pagination + 1
        }
    }
}