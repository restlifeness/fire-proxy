"""Add admin user from env

Revision ID: 04ee871e30c9
Revises: 534b7b1944c3
Create Date: 2023-07-01 19:37:59.845535

"""
from alembic import op
import sqlalchemy as sa
from os import getenv
from uuid import uuid4
from sqlalchemy.sql import table, column
from dotenv import load_dotenv
from werkzeug.security import generate_password_hash


# revision identifiers, used by Alembic.
revision = '04ee871e30c9'
down_revision = '534b7b1944c3'
branch_labels = None
depends_on = None


load_dotenv()


def upgrade() -> None:
    users = table('users',
        column('uuid', sa.String),
        column('username', sa.String),
        column('hashed_password', sa.Text),
        column('email', sa.String),
        column('is_active', sa.Boolean),
        column('is_admin', sa.Boolean)
    )

    hashed_password = generate_password_hash(getenv("ADMIN_PASSWORD"), method='scrypt')

    op.bulk_insert(users,
        [
            {
                "uuid": str(uuid4()),
                "username": getenv("ADMIN_USERNAME"),
                "hashed_password": hashed_password,
                "email": getenv("ADMIN_EMAIL"),
                "is_active": True,
                "is_admin": True
            }
        ]
    )


def downgrade() -> None:
    op.execute("DELETE FROM users WHERE username = '{}'".format(getenv("ADMIN_USERNAME")))
