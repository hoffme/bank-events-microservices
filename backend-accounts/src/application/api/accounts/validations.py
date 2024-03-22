from typing import Literal
from pydantic import BaseModel

class SearchAccountQueryDto(BaseModel):
    query: str = ""
    limit: str = 10
    skip: str = 0
    order_by: str = 'created_at'
    order_dir: str = 'desc'

class CreateAccountBodyDto(BaseModel):
    name: str = ""
    balance: int = 0
    currency: Literal['ARS', 'USD'] = "ARS"
    
class ModifiyAccountBodyDto(BaseModel):
    name: str = ""
