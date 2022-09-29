module.exports = [
    {
        message: "対象のentityは?　ex) User, Wallet　※パスカルケース",
        type: 'input',
        name: 'entity',
    },
    {
        message: "何をするuseCase?  ex) GetUserList, CreateUser UpdateUser　※パスカルケース",
        type: 'input',
        name: 'useCaseName',
    },
    {
        message: "useCaseで使われるレポジトリーのメソッド名?  ex) FindAll, FindByID, Create, Update, DeleteById　※パスカルケース",
        type: 'input',
        name: 'method',
    },
    {
        message: "outputのstruct名は?  ex) UserByID, UserList CreateBlock ※パスカルケース",
        type: 'input',
        name: 'output',
    },
    {
        message: "何を行うusecaseなのか説明を入力?  ex) トランザクションを作成する",
        type: 'input',
        name: 'desc',
    },
];
