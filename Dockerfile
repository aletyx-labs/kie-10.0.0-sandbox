# This Dockerfile is used to build a custom KIE Sandbox image with Aletyx branding.
# For community image use: apache/incubator-kie-sandbox-webapp:main
# For enterprise image use: aletyx-playground:10.0.0
FROM aletyx-playground:10.0.0

# Set working directory
WORKDIR /kie-sandbox/app

ENV KIE_SANDBOX_APP_NAME="aletyx playground"
ENV KIE_SANDBOX_FEEDBACK_URL="https://aletyx.ai/contact-us/"

# Copy branding files into the container
COPY branding/aletyx-light.svg images/app_logo_rgb_fullcolor_reverse.svg
COPY branding/aletyx-dark.svg images/app_logo_rgb_fullcolor_default.svg
COPY branding/favicon.svg favicon.svg

# Expose necessary ports
EXPOSE 8080

# Default command (if not already set in the base image, which it should from Community)
ENTRYPOINT [ "/kie-sandbox/entrypoint.sh" ]
