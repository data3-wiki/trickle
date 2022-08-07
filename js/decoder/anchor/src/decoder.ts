import { BorshAccountsCoder, ACCOUNT_DISCRIMINATOR_SIZE, Idl } from "@project-serum/anchor"; 

interface DecodedAccount {
    accountType: string
    decoded: any
}

export function decodeAccount(idl: Idl, value: string): DecodedAccount {
    // Construct map of unique 8 byte discriminator to account name.
    const discriminatorToAccountName = new Map();
    idl.accounts.forEach((acc) => {
        discriminatorToAccountName.set(BorshAccountsCoder.accountDiscriminator(acc.name).toString(), acc.name)
    });

    // Extract discriminator from account data, use it find account name, then use account name to decode.
    const buff = Buffer.from(value, "base64")
    const discriminator = buff.slice(0, ACCOUNT_DISCRIMINATOR_SIZE).toString()
    const accountName = discriminatorToAccountName.get(discriminator)
    const coder = new BorshAccountsCoder(idl)
    return {
        accountType: accountName,
        decoded: coder.decode(accountName, buff),
    }
}