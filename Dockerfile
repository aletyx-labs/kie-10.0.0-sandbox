# Use the existing community image as the base
FROM apache/incubator-kie-sandbox-webapp:10.0.0


# Set working directory
WORKDIR /kie-sandbox/app

#ENV KIE_SANDBOX_APP_NAME 'aletyx playground'

# Copy branding files into the container
COPY branding/aletyx-light.svg images/app_logo_rgb_fullcolor_reverse.svg
COPY branding/aletyx-dark.svg images/app_logo_rgb_fullcolor_default.svg
COPY branding/favicon.png favicon.png

# Expose necessary ports
EXPOSE 8080

# Default command (if not already set in the base image, which it should from Community)
ENTRYPOINT [ "/kie-sandbox/entrypoint.sh" ]
