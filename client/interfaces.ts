export interface Account {
    accountId: string;
    username: string;
    email: string;
    gender: 'male' | 'female' | 'other';
    age: number;
    money: number;
    image_uri: string;
}

export interface JWT {
    jwt: string
}
