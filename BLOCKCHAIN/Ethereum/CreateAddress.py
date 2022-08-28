from eth_account import Account
import secrets

priv = secrets.token_hex(32)
private_key = "0x" + priv
print(priv, 'private key is', "0x" + private_key)
acct = Account.from_key(private_key)

print("address", acct.address)

