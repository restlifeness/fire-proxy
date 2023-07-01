
from sqlalchemy import *
from sqlalchemy.orm import relationship
from sqlalchemy.ext.declarative import declarative_base


Base = declarative_base()


class BaseModel(Base):
    __abstract__ = True

    id = Column(Integer, primary_key=True)

    created_at = Column(DateTime, default=func.current_timestamp())
    updated_at = Column(
        DateTime, 
        default=func.current_timestamp(),
        onupdate=func.current_timestamp()
    )


class User(BaseModel):
    __tablename__ = 'users'

    uuid = Column(String(36), unique=True, nullable=False)
    username = Column(String(50), unique=True, nullable=False)
    password = Column(String(50), nullable=False)
    email = Column(String(50), unique=True, nullable=False)

    is_active = Column(Boolean, default=True)
    is_admin = Column(Boolean, default=False)

    def __repr__(self):
        return f'<User {self.username}>'


class Proxy(BaseModel):
    __tablename__ = 'proxy'

    address = Column(String(50), nullable=False)
    port = Column(Integer, nullable=False)

    stil_alive = Column(Boolean, default=True)

    def __repr__(self):
        return f'<ProxyConnection {self.user.username} {self.proxy.ip}>'
